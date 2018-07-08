import { rem } from "../../utils/style"

function renderItems(items, depth = 0) {
    return items.map(item => {
        const activity = item.isActive ? "active" : "inactive"
        return (
            <div
                key={item.slug}
                className={`widget-hierarchy-item${depth ? " widget-hierarchy-child" : ""}`}
            >
                <strong>{item.title}</strong>
                <span className="widget-hierarchy-detail">{item.template}</span>
                <span className={`widget-hierarchy-activity ${activity}`}>{activity}</span>
                {!!item.children.length && renderItems(item.children, depth + 1)}
            </div>
        )
    })
}

const Hierarchy = props => {
    return (
        <div className="hierarchy">
            {renderItems(props.children)}
            <style global jsx>{`
                .hierarchy > .widget-hierarchy-item {
                    border-bottom: 1px solid #f5f5f5;
                }

                :global(.widget-hierarchy-item) {
                    line-height: ${rem(50)};
                    font-size: ${rem(14)};
                    font-family: "Lato";
                    color: #091c3a;
                    font-weight: 300;
                    position: relative;
                }

                :global(.widget-hierarchy-item strong) {
                    color: #091c3a;
                    font-weight: 400;
                    display: inline-block;
                    margin-right: 30px;
                }

                :global(.widget-hierarchy-child) {
                    border-top: 1px solid #f5f5f5;
                    padding-left: 60px;
                }

                :global(.widget-hierarchy-child:before) {
                    content: "";
                    display: inline-block;
                    width: 30px;
                    border-bottom: 1px solid #091c3a;
                    position: absolute;
                    top: ${rem(25)};
                    left: 15px;
                }

                :global(.widget-hierarchy-activity) {
                    display: flex;
                    float: right;
                    align-items: center;
                }

                :global(.widget-hierarchy-activity:after) {
                    display: inline-block;
                    content: '';
                    width: 10px;
                    height: 10px;
                    background-color: #c0c0c0;
                    margin-left: 5px;
                    border-radius: 50%;
                }

                :global(.widget-hierarchy-activity.active:after) {
                    background-color: #b8e986;
                }


            `}</style>
        </div>
    )
}

export default Hierarchy
