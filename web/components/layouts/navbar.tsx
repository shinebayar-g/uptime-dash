import { HStack, Button, Grid, GridItem } from '@chakra-ui/react';

export default function Navbar() {
    return (
        <Grid
            templateColumns={'1fr 1fr'}
            templateRows={'48px'}
            templateAreas={`"left right"`}
            alignItems={'center'}
        >
            <GridItem area={'left'} ml={1}>
                <Button variant='ghost'>Dashboard</Button>
            </GridItem>
            <GridItem area={'right'}>
                <HStack spacing={1} justifyContent='right' mr={1}>
                    <Button variant='ghost'>Settings</Button>
                    <Button variant='ghost'>Sign in</Button>
                </HStack>
            </GridItem>
        </Grid>
    );
}
