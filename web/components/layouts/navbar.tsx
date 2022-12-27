import { HStack, Button, Grid, GridItem, Link } from '@chakra-ui/react';
import NextLink from 'next/link';

export default function Navbar() {
    return (
        <Grid
            templateColumns={'1fr 1fr'}
            templateRows={'48px'}
            templateAreas={`"left right"`}
            alignItems={'center'}
        >
            <GridItem area={'left'} ml={1}>
                <Link as={NextLink} href={`/`} style={{ textDecoration: 'none' }}>
                    <Button variant='ghost'>Dashboard</Button>
                </Link>
            </GridItem>
            <GridItem area={'right'}>
                <HStack spacing={1} justifyContent='right' mr={1}>
                    <Link as={NextLink} href={`/targets/add`} style={{ textDecoration: 'none' }}>
                        <Button variant='outline'>Add</Button>
                    </Link>
                    <Button variant='ghost'>Settings</Button>
                    <Button variant='ghost'>Sign in</Button>
                </HStack>
            </GridItem>
        </Grid>
    );
}
