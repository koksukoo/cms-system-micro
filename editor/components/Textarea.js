import { rem } from '../utils/style'

const Textarea = props => {
    const { label, value, ...rest } = props
    return (
        <div className="textarea">
            <label>{label}</label>
            <textarea {...rest}>{value}</textarea>
            <style jsx>{`
                .textarea {
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

                textarea {
                    flex: 1;
                }
            `}</style>
        </div>
    )
}

export default Textarea
