import { compose, withHandlers, withState, lifecycle } from 'recompose'
import { rem } from 'utils/style'

const enhance = compose(
    withState('value', 'setValue', ''),
    withHandlers({
        onValueChange: ({ setValue }) => (_, __, value) => setValue(value),
    }),
    lifecycle({
        componentDidMount() {
            if (!this.props.initialValue) return;
            this.setState({ value: this.props.initialValue})
        }
    })
)

const Input = props => {
    const { label, type, value, ...rest } = props
    return (
        <div className="input">
            <label>{label}</label>
            <input value={value} onChange={props.onValueChange} type={type} />
            <style jsx>{`
                .input {
                    display: flex;
                    flex-direction: row;
                    width: 100%;
                    margin: 5px 0;
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
                    flex: unset;
                    width: 40px;
                    height: 40px;
                    margin: 0;
                    appearance: none;
                    background-color: #fafafa;
                    border: 1px solid rgba(0, 0, 0, 0.1);
                    box-shadow: 0 1px 2px rgba(0,0,0,0.05), inset 0px -15px 10px -12px rgba(0,0,0,0.05);
                    padding: 9px;
                    border-radius: 5px;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    position: relative;
                }

                input[type="checkbox"]:active,
                input[type="checkbox"]:checked:active {
                    box-shadow: 0 1px 2px rgba(0,0,0,0.05), inset 0px 1px 3px rgba(0,0,0,0.1);
                }

                input[type="checkbox"]:checked {
                    border: 1px solid rgba(0,0,0,0.2);
                    box-shadow: 0 1px 2px rgba(0,0,0,0.05), inset 0px -15px 10px -12px rgba(0,0,0,0.05), inset 15px 10px -12px rgba(255,255,255,0.1);
                }

                input[type="checkbox"]:checked:after {
                    content: '\\2713';
                    font-size: ${rem(18)};
                    font-weight: 800;
                    color: rgba(0,0,0,0.6);
                }
            `}</style>
        </div>
    )
}

export default enhance(Input)
