import Head from 'next/head'
import Header from './Header'
import PageChooser from './PageChooser';

const Layout = (props) => (
    <div>
        <Head>
            <link href="https://fonts.googleapis.com/css?family=Lato:300,400" rel="stylesheet" />
        </Head>
        <Header />
        <PageChooser />

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
                max-width: 1200px;
                margin: 0 auto;
                padding: 0;
                font-family: 'Lato';
                -webkit-font-smoothing: antialiased;
                -moz-osx-font-smoothing: grayscale;
                text-rendering: geometricPrecision;
            }
            .content {
                padding: 20px 0;
            }
        `}</style>
    </div>
)

export default Layout