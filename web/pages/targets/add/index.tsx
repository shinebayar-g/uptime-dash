import { useEffect, useState } from 'react';
import NextLink from 'next/link';
import {
    Box,
    Button,
    FormControl,
    FormLabel,
    Grid,
    GridItem,
    Input,
    Link,
    NumberDecrementStepper,
    NumberIncrementStepper,
    NumberInput,
    NumberInputField,
    NumberInputStepper,
    Select,
    VStack,
} from '@chakra-ui/react';
import { createConnectTransport, createPromiseClient } from '@bufbuild/connect-web';
import { UptimeService } from '../../../gen/uptime_connectweb';
import { Target, Target_Type } from '../../../gen/uptime_pb';

const transport = createConnectTransport({
    baseUrl: 'http://localhost:8080',
});

const client = createPromiseClient(UptimeService, transport);

const types = (
    Object.keys(Target_Type).filter((s) => {
        return isNaN(Number(s));
    }) as (keyof typeof Target_Type)[]
).map((key, index) => {
    return Target_Type[index];
});

export default function AddTarget() {
    const [targets, setTargets] = useState<Target[]>([]);
    const [name, setName] = useState<string>('');
    const [targetType, setTargetType] = useState<Target_Type>(Target_Type.UNSPECIFIED);
    const [intervalSeconds, setIntervalSeconds] = useState<number>(20);
    const [timeoutSeconds, setTimeoutSeconds] = useState<number>(3);
    const [url, setUrl] = useState<string>('');
    const [hostname, setHostname] = useState<string>('');
    const [port, setPort] = useState<number>();

    const renderSwitch = (param: Target_Type) => {
        switch (param) {
            case Target_Type.HTTP:
                return (
                    <>
                        <FormLabel>URL</FormLabel>
                        <Input
                            placeholder='https://www.google.com'
                            type={'url'}
                            value={url}
                            onChange={(e) => setUrl(e.currentTarget.value)}
                        ></Input>
                    </>
                );
            case Target_Type.TCP:
                return (
                    <>
                        <FormLabel>Hostname / IP address</FormLabel>
                        <Input
                            placeholder='ec2-12-34-56-78.us-west-2.compute.amazonaws.com'
                            value={hostname}
                            onChange={(e) => setHostname(e.currentTarget.value)}
                        ></Input>

                        <FormLabel>Port</FormLabel>
                        <NumberInput min={1} max={65535} onChange={(val) => setPort(parseInt(val))}>
                            <NumberInputField placeholder={'8080'} />
                            <NumberInputStepper>
                                <NumberIncrementStepper />
                                <NumberDecrementStepper />
                            </NumberInputStepper>
                        </NumberInput>
                    </>
                );
            case Target_Type.PING:
                return (
                    <>
                        <FormLabel>Hostname / IP address</FormLabel>
                        <Input
                            placeholder='ec2-12-34-56-78.us-west-2.compute.amazonaws.com'
                            value={hostname}
                            onChange={(e) => setHostname(e.currentTarget.value)}
                        ></Input>
                    </>
                );
        }
    };

    const addTarget = async () => {
        const t = new Target({
            name,
            type: targetType,
            intervalSeconds,
            timeoutSeconds,
            url,
            hostname,
            port,
        });

        try {
            const res = await client.createTarget({ target: t });
            console.log(res);
        } catch (err) {
            console.error("couldn't create a new target.", err);
        }
    };

    useEffect(() => {
        client.getAllTargets({}).then((res) => {
            setTargets(res.targets);
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
                    {targets.map((target) => {
                        return (
                            <Link key={target.id} as={NextLink} href={`/targets/${target.id}`}>
                                {target.name}
                            </Link>
                        );
                    })}
                </VStack>
            </GridItem>
            <GridItem pl='2' area={'main'}>
                <FormControl>
                    <VStack alignItems={'left'}>
                        <FormLabel>Name</FormLabel>
                        <Input
                            type='email'
                            value={name}
                            onChange={(e) => setName(e.currentTarget.value)}
                        ></Input>

                        <FormLabel>Type</FormLabel>
                        <Select
                            placeholder='Select target type'
                            onChange={(e) => {
                                setTargetType(
                                    Target_Type[e.currentTarget.value as keyof typeof Target_Type]
                                );
                            }}
                        >
                            {types.map((t) => {
                                if (t === 'UNSPECIFIED') return;
                                return <option key={t}>{t}</option>;
                            })}
                        </Select>

                        {targetType && (
                            <>
                                <FormLabel>Interval Seconds</FormLabel>
                                <NumberInput
                                    min={15}
                                    value={intervalSeconds}
                                    onChange={(val) => setIntervalSeconds(parseInt(val))}
                                >
                                    <NumberInputField />
                                    <NumberInputStepper>
                                        <NumberIncrementStepper />
                                        <NumberDecrementStepper />
                                    </NumberInputStepper>
                                </NumberInput>

                                <FormLabel>Timeout Seconds</FormLabel>
                                <NumberInput
                                    min={1}
                                    max={10}
                                    value={timeoutSeconds}
                                    onChange={(val) => setTimeoutSeconds(parseInt(val))}
                                >
                                    <NumberInputField />
                                    <NumberInputStepper>
                                        <NumberIncrementStepper />
                                        <NumberDecrementStepper />
                                    </NumberInputStepper>
                                </NumberInput>
                            </>
                        )}

                        {targetType && renderSwitch(targetType)}

                        <Button onClick={addTarget}>Add</Button>
                    </VStack>
                </FormControl>
            </GridItem>
            <GridItem pl='2' bg='blue.300' area={'footer'}>
                Footer
            </GridItem>
        </Grid>
    );
}
