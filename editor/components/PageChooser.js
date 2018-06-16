import Link from 'next/link'
import { withRouter } from 'next/router';
import { rem } from '../utils/style';

function checkActive(route, match) {
    return route === match ? 'active' : ''
}

const PageChooser = withRouter((context) => {
    console.log(context.router.route);
    return (
    <div>
        <nav>
            <Link href="/">
                <a className={checkActive(context.router.route, '/')}>Pages</a>
            </Link>
            <Link href="/templates">
                <a className={checkActive(context.router.route, '/templates')}>Templates</a>
            </Link>
        </nav>
        <style jsx>{`
            height: 60px;
            display: flex;
            align-items: center;
            background-color: #091c3a;

            nav {
                padding: 0 15px;
            }

            a {
                color: #fff;
                text-decoration: none;
                margin-right: 20px;
                font-weight: 400;
                font-size: ${rem(15)};
                position: relative;
            }

            .active:after {
                content: '';
                width: 10px;
                height: 10px;
                background-color: #091c3a;
                position: absolute;
                bottom: -5px;
                left: calc(50% - 5px);
                transform: rotate(45deg);
            }
        `}</style>
    </div>
)})

export default PageChooser