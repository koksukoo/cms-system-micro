import Layout from 'components/PublicLayout'
import LoginForm from 'components/LoginForm'

const Login = props => (
    <Layout>
        <LoginForm />
        <style jsx>{`
            display: flex;
            justify-conten: center;
            align-items: center;
            height: 100%;
        `}</style>
    </Layout>
)

export default Login
