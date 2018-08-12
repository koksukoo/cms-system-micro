import { rem } from "utils/style"

const Input = props => {
    const { label, type, value, checked, hasError, ...rest } = props
    return (
        <div className="input">
            <label>{label}</label>
            <input
                value={value}
                className={hasError ? "error" : ""}
                type={type}
                checked={type === "checkbox" ? value : null}
                {...rest}
            />
            <style jsx>{`
                .input {
                    display: flex;
                    flex-direction: row;
                    width: 100%;
                    margin: 5px 0;
                }

                .error {
                    border-color: #ff4d4d;
                    box-shadow: 0 0 5px 0 #ff9f9f;
                }

                label {
                    display: inline-block;
                    width: 20%;
                    text-align: right;
                    line-height: ${rem(30)};
                    padding: 5px 10px;
                }

                input {
                    flex: 1;
                    border: 1px solid rgba(0, 0, 0, 0.1);
                    border-radius: 5px;
                    font-size: ${rem(16)};
                    padding: 0 10px;
                }

                input[type="checkbox"] {
                    cursor: pointer;
                    flex: unset;
                    width: 60px;
                    height: 30px;
                    margin: 5px 0 0;
                    appearance: none;
                    background-color: #fff;
                    border: 1px solid rgba(0, 0, 0, 0.1);
                    padding: 0;
                    border-radius: 20px;
                    display: flex;
                    justify-content: flex-start;
                    align-items: center;
                    position: relative;
                    transition: background-color 0.2s;
                }

                input[type="checkbox"]:after {
                    content: "";
                    display: inline-block;
                    width: 26px;
                    height: 26px;
                    background-color: #ececec;
                    border-radius: 50%;
                    margin: 0 2px;
                    transition: background-color 0.2s;
                }

                input[type="checkbox"]:checked {
                    justify-content: flex-end;
                }

                input[type="checkbox"]:checked:after {
                    background-color: #184080;
                }

                input[type="checkbox"]:checked:before {
                    content: '\u2713';
                    position: absolute;
                    font-weight: 800;
                    left: 10px;
                    color: rgba(0, 0, 0, 0.2);
                }
            `}</style>
        </div>
    )
}

export default Input
