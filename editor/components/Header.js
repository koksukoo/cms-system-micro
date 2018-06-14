import Link from 'next/link'

const Header = () => (
    <div>
        <Link href="/">
            <a>Home</a>
        </Link>
        <Link href="/edit">
            <a>Edit</a>
        </Link>
        <style jsx>{`
            a {
                font-family: "Arial";
                margin-right: 20px;
                text-decoration: none;
                color: blue;
            }

            a:hover {
                opacity: 0.6;
            }
            `}</style>
    </div>
)

export default Header