import { createConnectTransport, createPromiseClient, ConnectError } from '@bufbuild/connect-web';
import { Box, Grid, GridItem, Link, VStack } from '@chakra-ui/react';
import { GetServerSideProps } from 'next';
import NextLink from 'next/link';
import { useEffect, useState } from 'react';
import { UptimeService } from '../../../gen/uptime_connectweb';
import { GetTargetResponse, Target } from '../../../gen/uptime_pb';

const transport = createConnectTransport({
    baseUrl: 'http://localhost:8080',
});

const client = createPromiseClient(UptimeService, transport);

export const getServerSideProps: GetServerSideProps = async (context) => {
    let res: GetTargetResponse | undefined;
    try {
        // @ts-ignore
        res = await client.getTarget({ id: context.params.id });
    } catch (err) {
        if (err instanceof ConnectError) {
            if (err.code === 5) {
                return {
                    notFound: true,
                };
            } else {
                console.log('unknown connect error,', err);
            }
        } else {
            console.error('unknown error.', err);
        }
    }
    return {
        props: {
            target: res?.target?.toJson(),
        },
    };
};

export default function TargetDetail({ target }: { target: Target }) {
    const [data, setData] = useState<Target[]>([]);

    useEffect(() => {
        client.getAllTargets({}).then((res) => {
            setData(res.targets);
        });
    }, []);

    return (
        <Grid
            templateColumns={'1fr 3fr'}
            templateRows={'4fr minmax(50px, 1fr)'}
            templateAreas={`"side main"
                            "side footer"`}
            gap='1'
            color='blackAlpha.700'
            fontWeight='bold'
            height='calc(100vh - 48px)'
        >
            <GridItem pl='2' bg='pink.300' area={'side'}>
                <VStack align='stretch'>
                    <Box>Targets</Box>
                    {data.map((target) => {
                        return (
                            <Link key={target.id} as={NextLink} href={`/targets/${target.id}`}>
                                {target.name}
                            </Link>
                        );
                    })}
                </VStack>
            </GridItem>
            <GridItem pl='2' bg='green.300' area={'main'}>
                <Box>ID: {target.id}</Box>
                <Box>Name: {target.name}</Box>
                <Box>Type: {target.type}</Box>
                <Box>Interval Seconds: {target.intervalSeconds}</Box>
                <Box>Timeout Seconds: {target.timeoutSeconds}</Box>
                <Box>Url: {target.url}</Box>
                <Box>Hostname: {target.hostname}</Box>
                <Box>Port: {target.port}</Box>
            </GridItem>
            <GridItem pl='2' bg='blue.300' area={'footer'}>
                Footer
            </GridItem>
        </Grid>
    );
}
