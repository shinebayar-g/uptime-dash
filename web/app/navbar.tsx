'use client';

import { useColorModeValue, Flex, HStack, Button, chakra } from '@chakra-ui/react';

export default function Navbar() {
    const bg = useColorModeValue('white', 'gray.800');
    return (
        <>
            <chakra.header
                bg={bg}
                w='full'
                px={{
                    base: 2,
                    sm: 4,
                }}
                py={3}
                shadow='md'
            >
                <Flex alignItems='center' justifyContent='space-between' mx='auto'>
                    <Flex>
                        <Button w='full' variant='ghost'>
                            Dashboard
                        </Button>
                    </Flex>

                    <HStack display='flex' alignItems='center' spacing={1}>
                        <HStack
                            spacing={1}
                            mr={1}
                            color='brand.500'
                            display={{
                                base: 'none',
                                md: 'inline-flex',
                            }}
                        >
                            <Button variant='ghost'>Settings</Button>
                            <Button variant='ghost'>Sign in</Button>
                        </HStack>
                    </HStack>
                </Flex>
            </chakra.header>
        </>
    );
}
