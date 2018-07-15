const ActionButton = (props) => (
    <button {...props} className="button">
        {props.children}
        <style jsx>{`
            .button {
                cursor: pointer;
                padding: 10px 15px;
                color: #091c3a;
                background-color: #fff;
                border: 0;
                border-bottom: 2px solid transparent;
                font-size: 0.9rem;
            }
            .button:hover {
                color: #184080;
                background-color: #f6f6f6;
                border-bottom: 2px solid #184080;
            }

            :global(.button > svg) {
                width: 1.1rem;
                height: 1.1rem;
                padding-right: 5px;
                color: #091c3a;
            }

            :global(.button:hover svg) {
                color: #184080;
            }
        `}</style>
    </button>
)

export default ActionButton