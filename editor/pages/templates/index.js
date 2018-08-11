import Link from "next/link"
import PlusIcon from "react-icons/lib/md/add"
import Layout from "components/BaseLayout"
import Widget from "components/Widget"
import Button from "components/ActionButton"
import List from "components/widgets/List"
import { fetchTemplates } from "io"

const Templates = props => {
    return (
        <Layout>
            <Widget title="Template list">
                <Widget.Actions>
                    <Link href="/templates/new">
                        <Button>
                            <PlusIcon /> New Template
                        </Button>
                    </Link>
                </Widget.Actions>
                <Widget.Body>
                    <List children={props.templates} />
                </Widget.Body>
            </Widget>
        </Layout>
    )
}

Templates.getInitialProps = async () => {
    const res = await fetchTemplates()
    const data = await res.json()

    return await {
        templates: (!data) ? [] : data.map(t => {
            t.url = `/templates/edit`
            return t;
        })
    }
}

export default Templates
