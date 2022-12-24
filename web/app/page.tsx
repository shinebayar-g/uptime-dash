import { createConnectTransport, createPromiseClient } from '@bufbuild/connect-web'
import { UptimeService } from '../gen/uptime_connectweb'

const transport = createConnectTransport({
    baseUrl: "http://localhost:8080",
})

const client = createPromiseClient(UptimeService, transport)

export default function Home() {
    return (
        <h1>Hello World</h1>
    )
}
