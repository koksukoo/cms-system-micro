import Link from 'next/link';
import { rem } from 'utils/style';

const Widget = (props) => {
    const { title, children, ancestors, ...rest } = props;

    return (
        <section className="widget">
            <div className="wh">
                <h2>{ancestors &&
                    ancestors.map(a => <Link key={a.href} href={a.href}><a>{a.title}</a></Link>)}
                    { title }
                </h2>
                <Widget.Actions { ...{children: children.slice(0, 1)} } />
            </div>
            
            <Widget.Body {...{children: children.slice(1)} } {...rest} />

            <style jsx>{`
                .widget {
                    box-shadow: 0 1px 9px 0 rgba(0,0,0,0.08);
                    border-radius: 2px;
                    border: 1px solid rgba(0,0,0,0.1);
                    padding: ${rem(10)} ${rem(20)};
                }
                .wh {
                    display: flex;
                    justify-content: space-between;
                    align-items: center;
                    position: relative;
                }
                .wh:after {
                    display: block;
                    content: '';
                    border-bottom: 1px solid #e4e4e4;
                    position: absolute;
                    bottom: 0;
                    left: -${rem(20)};
                    right: -${rem(20)};
                }
                h2 {
                    color: #4A4A4A;
                    font-weight: 400;
                    margin: 10px 0 20px;
                }
                h2 a {
                    color: rgba(0, 0, 0, 0.4);
                    text-decoration: none;
                    padding-right: 5px;
                }
                h2 a:after {
                    content: ' /';
                }
            `}</style>
        </section>
    )
}

Widget.Actions = (props) => (
    <div>{props.children}</div>
)

Widget.Body = (props) => (
    <div>{props.children}</div>
)

export default Widget;