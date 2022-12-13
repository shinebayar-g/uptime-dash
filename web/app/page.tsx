'use client';
import { useState, useEffect } from 'react'
import { createConnectTransport, createPromiseClient } from '@bufbuild/connect-web'

import Image from 'next/image'
import styles from './page.module.css'
import { UptimeService } from '../gen/uptime_connectweb'

const transport = createConnectTransport({
    baseUrl: "http://localhost:8080",
})

const client = createPromiseClient(UptimeService, transport)

export default function Home() {
    const [data, setData] = useState(null)
    const [isLoading, setLoading] = useState(false)

    useEffect(() => {
        setLoading(true)
        client.getAllTargets({})
            .then((res) => res.toJSON())
            .then((data: any) => {
                setData(data)
                setLoading(false)
            })
    }, [])

    if (isLoading) return <p>Loading...</p>
    if (!data) return <p>No profile data</p>

    return (
        <div>
            <main className={styles.main}>
                <div>targets:
                    {JSON.stringify(data)}
                </div>
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
