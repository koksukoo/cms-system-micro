import { rem } from '../../utils/style'

function renderItems(items, depth = 0) {
    return items.map(item => (
        <div key={item.slug} className={`widget-hierarchy-item${depth ? ' widget-hierarchy-child' : ''}`}>
            <strong>{item.title}</strong>
            {!!item.children.length && renderItems(item.children, depth + 1)}
        </div>
    ))
}



const Hierarchy = props => {
    console.log(props)
    return (
        <div className="hierarchy">
            {renderItems(props.children)}
            <style global jsx>{`
                .hierarchy > .widget-hierarchy-item{
                    border-bottom: 1px solid #f5f5f5;
                }

                :global(.widget-hierarchy-item) {
                    line-height: ${rem(50)}; 
                    font-size: ${rem(14)};
                    font-family: 'Lato';
                    color: #091C3A;
                    font-weight: 300;
                    position: relative;
                }

                :global(.widget-hierarchy-item strong) {
                    color: #091C3A;
                    font-weight: 400;
                }

                :global(.widget-hierarchy-child) {
                    border-top: 1px solid #f5f5f5;
                    padding-left: 60px;
                }

                :global(.widget-hierarchy-child:before) {
                    content: '';
                    display: inline-block;
                    width: 30px;
                    border-bottom: 1px solid #091C3A;
                    position: absolute;
                    top: ${rem(25)};
                    left: 15px;
                }
            `}</style>
        </div>
    )
}

export default Hierarchy
