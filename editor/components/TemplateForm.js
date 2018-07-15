import { withFormik, Form } from "formik"
import SaveIcon from "react-icons/lib/fa/floppy-o"
import DeleteIcon from "react-icons/lib/fa/trash-o"
import Router from "next/router"
import Input from "components/Input"
import FieldsInput from "components/FieldsInput"
import Codearea from "components/Codearea"
import Widget from "components/Widget"
import Button from "components/ActionButton"
import { createTemplate, updateTemplate, deleteTemplate } from "io"

const handleDeleteTemplate = event => {
    event.preventDefault()
    deleteTemplate(id) // todo !!!
}

const TemplateForm = ({
    ancestors,
    values,
    errors,
    touched,
    handleChange,
    handleBlur,
    isSubmitting,
    setFieldValue,
    setFieldTouched,
    title,
    type
}) => (
    <Form>
        <Widget title={title} ancestors={ancestors}>
            <Widget.Actions>
                {type === "edit" && (
                    <Button onClick={handleDeleteTemplate}>
                        <DeleteIcon /> Delete
                    </Button>
                )}
                <Button type="submit">
                    <SaveIcon /> Save template
                </Button>
            </Widget.Actions>
            <Widget.Body>
                <br />
                <Input
                    name="title"
                    label="Title"
                    onChange={handleChange}
                    onBlur={handleBlur}
                    value={values.title}
                    hasError={touched.title && errors.title}
                />
                <Input
                    name="isActive"
                    label="Active"
                    type="checkbox"
                    onChange={handleChange}
                    onBlur={handleBlur}
                    value={values.isActive}
                />
                <FieldsInput
                    name="fields"
                    label="Fields"
                    initialValue={values.fields ? values.fields : null}
                />
                <Codearea
                    name="content"
                    label="Content"
                    rows="20"
                    onChange={setFieldValue}
                    setFieldTouched={setFieldTouched}
                    onBlur={handleBlur}
                    value={values.content}
                    initialValue={values.content ? values.content : null}
                />
                <br />
            </Widget.Body>
        </Widget>
    </Form>
)

export default withFormik({
    mapPropsToValues: ({ values: { title, isActive, content } }) => ({ title, isActive, content }),
    validate(values, props) {
        const err = {}
        if (!values.title) {
            err.title = "Required"
        }
        return err
    },
    async handleSubmit(values, { props, setSubmitting, setErrors }) {
        switch (props.type) {
            case "create":
                await createTemplate(values)
                break
            case "edit":
                await updateTemplate(props.values.slug, values)
                break
        }

        setSubmitting(false)
        Router.push("/templates")
    }
})(TemplateForm)
