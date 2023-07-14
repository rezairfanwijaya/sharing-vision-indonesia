import { DOMAIN, SUCCESS } from '../constant/index'
import { useState, useEffect } from 'react'
import Swal from 'sweetalert2'
import withReactContent from 'sweetalert2-react-content'
import { useNavigate } from "react-router-dom";
import { Column } from 'primereact/column';
import { DataTable } from 'primereact/datatable';

const Preview = () => {
    const [articles, setArticles] = useState([])
    const [articleToShow, setArticleToShow] = useState([])
    const limit = 100
    const MySwal = withReactContent(Swal)
    const navigate = useNavigate();

    const getAllArticles = () => {
        fetch(DOMAIN + `/articles/${limit}/0`, {
            method: "GET"
        })
            .then(res => { return res.json() })
            .then(data => {
                if (data.meta.status === SUCCESS) {
                    setArticles(data.data.articles)
                    setArticleToShow(data.data.articles)
                } else {
                    MySwal.fire({
                        position: 'center',
                        icon: 'warning',
                        title: 'Gagal',
                        text: "terjadi kesalahan",
                        showConfirmButton: true,
                    })
                }
            })
    }

    // filter untuk tab publish, draft thrash
    const articleFilter = (status) => {
        const filteredArticles = articles.filter(article => article.status === status);
        setArticleToShow(filteredArticles)
    }

    
    useEffect(() => {
        getAllArticles()
        articleFilter("publish")
    }, [])

    useEffect(() => {
        articleFilter("publish")
    }, [articles])


    return (<>
        <div className="conten flex flex-col gap-3">
        <DataTable
                value={articleToShow}
                emptyMessage="Tidak ada data articles"
                currentPageReportTemplate="Showing {first} to {last} of {totalRecords} alumni"
                dataKey="id"
                paginator
                rows={limit}
                rowsPerPageOptions={[5, 10]}
                paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
            >
                <Column headerStyle={{ width: '3rem' }} exportable={false}></Column>
                <Column field="title" header="Title" style={{ width: '15%' }} ></Column>
                <Column field="content" header="Content" style={{ width: '75%' }} ></Column>
                <Column field="category" header="Category" style={{ width: '50%' }} ></Column>
            </DataTable>
        </div>
    </>);
}

export default Preview;