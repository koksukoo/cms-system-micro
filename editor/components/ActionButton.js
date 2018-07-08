const ActionButton = (props) => (
    <a {...props} className="button">
        {props.children}
        <style jsx>{`
            & {
                cursor: pointer;
                padding: 10px 20px;
                transition: background-color 0.2s ease-in-out;
            }
            .button:hover {
                background-color: #eaeaea;
            }
        `}</style>
    </a>
)

export default ActionButton