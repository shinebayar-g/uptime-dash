import { Inter } from '@next/font/google';
import Provider from './provider';
import Navbar from './navbar';

const inter = Inter({ subsets: ['latin'] });

export default function RootLayout({ children }: { children: React.ReactNode }) {
    return (
        <html lang='en' className={inter.className}>
            <head />
            <body>
                <Provider>
                    <Navbar />
                    {children}
                </Provider>
            </body>
        </html>
    );
}
