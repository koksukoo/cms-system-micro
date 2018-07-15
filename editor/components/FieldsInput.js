
import { compose, withHandlers, withState, lifecycle } from 'recompose'
import { connect } from 'formik'
import Button from './GeneralButton'
import { rem } from 'utils/style'

const enhance = compose(
    withState('value', 'setValue', ''),
    withHandlers({
        onValueChange: ({ setValue }) => (value) => {
            setValue(value)
        },
    }),
    lifecycle({
        componentDidMount() {
            if (!this.props.initialValue || this.props.value) return;
            this.props.onValueChange(this.props.initialValue)
        },
    })
)

const Fields = props => {
    const { label, value, name, formik} = props
    return (
        <div className="fields">
            <label>Fields</label>
            <input name="fields" type="hidden" value={[]} />

            <input type="text" placeholder="DOM selector" />
            <div className="select-wrapper">
                <select>
                    <option>Select field type</option>
                    <option>Normal</option>
                </select>
            </div>
            <Button>Add</Button>

            <style jsx>{`
                .fields {
                    display: flex;
                    flex-direction: row;
                    width: 100%;
                    margin: 5px 0 15px 0;
                }

                input[type="text"] {
                    flex: 2;
                    border: 1px solid rgba(0, 0, 0, 0.1);
                    border-radius: 5px;
                    font-size: ${rem(16)};
                    padding: 0 10px;
                }

                select {
                    appearance: button;
                    margin: 0 5px;
                    height: 42px;
                    background-color: transparent;
                    border: 1px solid rgba(0, 0, 0, 0.1);
                    border-radius: 5px;
                    padding: 0 25px 0 10px;;
                    font-size: ${rem(16)};
                    flex: 1;
                    cursor: pointer;
                }

                .select-wrapper {
                    position: relative;
                    display: flex;
                    flex: 1;
                }

                .select-wrapper:after {
                    content: '';
                    background-color: rgba(0,0,0,0.1);
                    height: 15px;
                    width: 15px;
                    transform: rotate(45deg);
                    position: absolute;
                    right: 10px;
                    top: 13px;
                    user-select: none;
                    z-index: -2;
                }

                .select-wrapper:before {
                    content: '';
                    position: absolute;
                    top: 19px;
                    right: 6px;
                    width: 22px;
                    border-bottom: 3px solid #fff;
                    display: inline-block;
                    z-index: -1;
                }

                select::after {
                    content: 's';
                    display: block;
                }

                label {
                    display: inline-block;
                    width: 20%;
                    text-align: right;
                    line-height: ${rem(30)};
                    padding: 5px 10px;
                }
            `}</style>
        </div>
    )
}
export default connect(enhance(Fields))