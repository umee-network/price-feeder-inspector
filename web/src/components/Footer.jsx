export default function Footer() {
    return (
        <div className="container">
            <div className="root-bg ubuntu-light py-3 my-4 border-top">
                <p>RPC: <a target="_blank" className="text-black no-anchor" href="https://umee-rpc.polkachu.com/">https://umee-rpc.polkachu.com</a></p>
                <p>LCD: <a target="_blank" className="text-black no-anchor" href="https://umee-api.polkachu.com/">https://umee-api.polkachu.com</a></p>
            </div>
            <div className="root-bg ubuntu-light py-3 my-4">
            Created by the <img src="https://app.ux.xyz/favicon.png" width={16}></img> Umee Team &middot; &copy; 2024
            </div>
        </div>
    )
}