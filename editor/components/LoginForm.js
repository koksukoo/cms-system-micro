import { withFormik, Form } from "formik"
import Router from "next/router"
import SubmitIcon from "react-icons/lib/fa/sign-in"
import Input from "components/Input"
import Button from "components/ActionButton"
import Hint from "components/Hint"
import Widget from "components/Widget"
import { login } from "io"

const LoginForm = ({
    errors,
    touched,
    handleChange,
    handleBlur,
    isSubmitting,
}) => (
    <Form>
        <Widget title="Log in, please...">
            <Widget.Actions>
                <Button type="submit">
                    <SubmitIcon /> Log in
                </Button>
            </Widget.Actions>
            <Widget.Body>
                <br />
                <Input
                    name="username"
                    label="Email"
                    onChange={handleChange}
                    onBlur={handleBlur}
                    hasError={touched.username && errors.username}
                />
                <Input
                    name="password"
                    label="Password"
                    onChange={handleChange}
                    onBlur={handleBlur}
                    hasError={touched.password && errors.password}
                />
                <br />
                <Hint>
                    Don't have an account? <a href="#">Register</a> | Forgot your password?{" "}
                    <a href="#">Reset password</a>
                </Hint>
            </Widget.Body>
        </Widget>
        <style jsx>{`
            min-height: 100vh;
        `}</style>
    </Form>
)

export default withFormik({
    mapPropsToValues: props => ({ username: '', password: '' }),
    validate(values, props) {
        const err = {}
        if (!values.username) {
            err.username = "Required"
        }
        if (!values.password) {
            err.password = "Required"
        }
        return err
    },
    async handleSubmit(values, { props, setSubmitting, setErrors }) {
        const res = await login(values)
        console.log(res);
        setSubmitting(false)
        Router.push("/")
    },
})(LoginForm)
