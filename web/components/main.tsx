import { useEffect, useState } from 'react';
import NextLink from 'next/link';
import { Box, Grid, GridItem, Link, VStack } from '@chakra-ui/react';
import { createConnectTransport, createPromiseClient } from '@bufbuild/connect-web';
import { UptimeService } from '../gen/uptime_connectweb';
import { Target } from '../gen/uptime_pb';

const transport = createConnectTransport({
    baseUrl: 'http://localhost:8080',
});

const client = createPromiseClient(UptimeService, transport);

export default function Main() {
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
                Main
            </GridItem>
            <GridItem pl='2' bg='blue.300' area={'footer'}>
                Footer
            </GridItem>
        </Grid>
    );
}
