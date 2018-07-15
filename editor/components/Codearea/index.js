import { Controlled as Codemirror } from 'react-codemirror2'
import { compose, withHandlers, withState, lifecycle } from 'recompose'
import { connect } from 'formik'
import { rem } from 'utils/style'

if (process.browser) {
    require('codemirror/mode/xml/xml');
}

const enhance = compose(
    withState('value', 'setValue', ''),
    withHandlers({
        onValueChange: ({ setValue }) => (_, __, value) => {
            setValue(value)
        },
    }),
    lifecycle({
        componentDidMount() {
            if (!this.props.initialValue || this.props.value) return;
            this.props.onValueChange(null, null, this.props.initialValue)
        },
    })
)

const Codearea = props => {
    const { label, value, name, formik} = props
    return (
        <div className="codearea">
            <label>{label}</label>
            <Codemirror
                value={value}
                onBeforeChange={props.onValueChange}
                onChange={() => formik.setFieldValue(name, value, false)}
                options={{
                    mode: 'xml',
                    htmlMode: true,
                    matchClosing: true,
                    theme: 'default',
                    lineNumbers: true,
                }}
            />
            <style jsx>{`
                .codearea {
                    display: flex;
                    flex-direction: row;
                }

                label {
                    display: inline-block;
                    width: 20%;
                    text-align: right;
                    line-height: ${rem(30)};
                    padding: 5px 10px;
                }

                :global(.react-codemirror2) {
                    flex: 1;
                    max-width: 80%;
                }

                :global(.cm-s-default) {
                    border: 1px solid rgba(0, 0, 0, 0.1);
                    border-radius: 5px;
                }
            `}</style>
        </div>
    )
}

export default connect(enhance(Codearea))
