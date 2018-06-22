import Link from 'next/link'
import Layout from '../components/BaseLayout'
import Widget from '../components/Widget'
import HierarchyBody from '../components/widgets/Hierarchy'
import { fetchPageHierarchy } from '../io';

const Index = (props) => (
    <Layout>
        <Widget title="Site map">
            <Widget.Actions>
                <Link href="/pages/new">
                    <a>+ New page</a>
                </Link>
            </Widget.Actions>

            <Widget.Body>
                <HierarchyBody children={props.hierarchy} />
            </Widget.Body>
        </Widget>
    </Layout>
)

Index.getInitialProps = async () => {
    const res = await fetchPageHierarchy()
    const data = await res.json()
    
    return await {
        hierarchy: data,
    }
}

export default Index