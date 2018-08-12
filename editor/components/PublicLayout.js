import Head from 'next/head'

const Layout = (props) => (
    <div>
        <Head>
            <link href="https://fonts.googleapis.com/css?family=Lato:300,400" rel="stylesheet" />
        </Head>

        <section className="content">
            {props.children}
        </section>

        <style jsx global>{`
            *:focus {
                outline: none;
            }
        `}</style>
        <style jsx>{`
            & {
                max-width: 600px;
                min-height: 100vh;
                margin: auto;
                padding: 0;
                font-family: 'Lato';
                -webkit-font-smoothing: antialiased;
                -moz-osx-font-smoothing: grayscale;
                text-rendering: geometricPrecision;
            }
            .content {
                padding: 40px 20px;
            }
        `}</style>
    </div>
)

export default Layout