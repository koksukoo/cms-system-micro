const ErrorMessage = ({children, ...rest}) => (
    <div {...rest} className="error">
        {children}
        <style jsx>{`
            .error {
                background: linear-gradient(to bottom, rgba(169,3,41,1) 0%, rgba(143,2,34,1) 44%, rgba(109,0,25,1) 100%);
                color: #fff;
                font-weight: bold;
                margin: 60px auto;
                text-align: center;
                padding: 30px 15px;
                width: 90%;
            }
        `}</style>
    </div>
)

export default ErrorMessage