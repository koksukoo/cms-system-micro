const GeneralButton = (props) => (
    <button {...props} className="button">
        {props.children}
        <style jsx>{`
            .button {
                cursor: pointer;
                padding: 10px 15px;
                font-size: 0.7rem;
                border: 0;
                background-color: #f1f1f1;
                border-bottom: 3px solid rgba(0,0,0,0.1);
                border-radius: 5px;
                font-weight: 800;
                color: #4c4b4b;
                text-transform: uppercase;
            }
            .button:hover {
                background-color: #f6f6f6;
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

export default GeneralButton