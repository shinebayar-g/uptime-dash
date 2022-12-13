import Image from 'next/image'
import styles from './page.module.css'
import { createConnectTransport, createPromiseClient } from '@bufbuild/connect-web'
import { UptimeService } from '../proto/gen/proto/uptime_connectweb'

const transport = createConnectTransport({
    baseUrl: "http://localhost:8080",
})

const client = createPromiseClient(UptimeService, transport)

export default function Home() {
    return (
        <div>
            <main className={styles.main}>
                <h1>Hello World!</h1>
            </main>

            <footer className={styles.footer}>
                <a
                    href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    Powered by{' '}
                    <span className={styles.logo}>
                        <Image src="/vercel.svg" alt="Vercel Logo" width={72} height={16} />
                    </span>
                </a>
            </footer>
        </div>
    )
}
