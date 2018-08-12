import { rem } from 'utils/style'

const Hint = props => {
    return (
        <p className="hint">
            {props.children}
            <style jsx>{`
                .hint {
                    border-top: 1px solid #f4f4f4;
                    font-size: ${rem(12)};
                    padding: 30px 0;
                    text-align: center;
                }
            `}</style>
        </p>
    )
}

export default Hint
