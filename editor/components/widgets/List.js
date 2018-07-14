import Link from "next/link"
import { rem } from "utils/style"

const List = props => {
    
    return (
        <div>
            {props && props.children.map(child => (
                <Link as={`${child.url}/${child.slug}`} href={`${child.url}?id=${child.slug}`} key={child.slug}>
                    <a className="widget-list-item">
                        <strong>{child.title}</strong>
                        <style jsx>{`
                        .widget-list-item {
                            border-bottom: 1px solid #f5f5f5;
                            display: block;
                            line-height: ${rem(50)};
                            font-size: ${rem(14)};
                            font-family: "Lato";
                            color: #091c3a;
                            font-weight: 300;
                            position: relative;
                            text-decoration: none;
                        }
                    `}</style>
                    </a>
                </Link>
            ))}
        </div>
    )
}

export default List
