import Head from 'next/head';
import { Inter } from '@next/font/google';
import Main from '../components/main';

const inter = Inter({ subsets: ['latin'] });

export default function Home() {
    return (
        <>
            <Head>
                <title>Uptime Dash - Free, Open Source uptime monitoring</title>
                <meta name='viewport' content='width=device-width, initial-scale=1' />
                <meta
                    name='description'
                    content='Uptime Dash - Free, Open Source uptime monitoring'
                />
                <link rel='icon' href='/favicon.ico' />
            </Head>
            <Main />
        </>
    );
}
