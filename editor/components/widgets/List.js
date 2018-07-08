import Link from "next/link"
import { rem } from "~/utils/style"

const List = props => {
    console.log(props.children)
    return (
        <div>
            {props.children.map(child => (
                <Link href={`/templates/edit/${child.slug}`}>
                    <a key={child.slug} className="widget-list-item">
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
