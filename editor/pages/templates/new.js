import { Component } from 'react'
import Layout from '../../components/BaseLayout'
import Widget from '../../components/Widget'
import Button from '../../components/ActionButton'
import Input from '../../components/Input'
import Codearea from '../../components/Codearea'
import { createTemplate } from '../../io'

export default class NewTemplate extends Component {
    constructor(props) {
        super(props)
        this.handleForm = this.handleForm.bind(this)
    }

    async handleForm() {
        const {title, isActive, content} = this.form
        const data = {
            title: title.value,
            isActive: isActive.checked,
            content: content.value,
        }

        const res = await createTemplate(data)
        const json = await res.json()
        console.log(json)
    }

    render() {
        return (
            <Layout>
                <Widget title="New Template">
                    <Widget.Actions>
                        <Button onClick={this.handleForm}>Save</Button>
                    </Widget.Actions>
                    <Widget.Body>
                        <form onSubmit={createTemplate} ref={node => (this.form = node)}>
                            <Input name="title" label="Title" required />
                            <Input name="isActive" label="Is active" type="checkbox" />
                            <Codearea name="content" label="Content" rows="20" />
                        </form>
                    </Widget.Body>
                </Widget>
                <style jsx>{`
                    form {
                        margin: 30px 0;
                    }
                `}</style>
            </Layout>
        )
    }
}