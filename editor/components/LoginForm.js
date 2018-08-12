import { withFormik, Form } from "formik"
import Router from "next/router"
import SubmitIcon from "react-icons/lib/fa/sign-in"
import Input from "components/Input"
import Button from "components/ActionButton"
import Hint from "components/Hint"
import Widget from "components/Widget"

const LoginForm = ({
    errors,
    touched,
    handleChange,
    handleBlur,
    isSubmitting,
    setFieldValue,
    setFieldTouched
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
                    hasError={touched.title && errors.title}
                />
                <Input
                    name="password"
                    label="Password"
                    onChange={handleChange}
                    onBlur={handleBlur}
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
    validate(values, props) {
        const err = {}
        if (!values.title) {
            err.title = "Required"
        }
        return err
    },
    async handleSubmit(values, { props, setSubmitting, setErrors }) {
        setSubmitting(false)
        // Router.push("/")
    }
})(LoginForm)
