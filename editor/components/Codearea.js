import {Controlled as Codemirror } from 'react-codemirror2'
import { compose, withHandlers, withState } from 'recompose';
import 'codemirror/lib/codemirror.css'
import { rem } from '../utils/style'

if (process.browser) {
    require('codemirror/mode/xml/xml');
}

const enhance = compose(
    withState('value', 'setValue', ''),
    withHandlers({
        onValueChange: ({ setValue }) => (_, __, value) => setValue(value),
    })
)

const Codearea = props => {
    const { label, value, name} = props
    return (
        <div className="codearea">
            <label>{label}</label>
            <input type="hidden" name={name} value={value} />
            <Codemirror
                value={value}
                onBeforeChange={props.onValueChange}
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
                }

                :global(.cm-s-default) {
                    border: 1px solid rgba(0, 0, 0, 0.1);
                    border-radius: 5px;
                }
            `}</style>
        </div>
    )
}

export default enhance(Codearea)
