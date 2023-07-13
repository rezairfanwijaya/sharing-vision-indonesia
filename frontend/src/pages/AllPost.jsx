import { DOMAIN, SUCCESS } from '../constant/index'
import { useState, useEffect } from 'react'
import Swal from 'sweetalert2'
import withReactContent from 'sweetalert2-react-content'
import { Column } from 'primereact/column';
import { Button } from 'primereact/button';
import { DataTable } from 'primereact/datatable';
import { Toolbar } from 'primereact/toolbar';
import { useNavigate } from "react-router-dom";


const AllPosts = () => {
    const [articles, setArticles] = useState([])
    const [articleToShow, setArticleToShow] = useState([])
    const [IsMoveToTrash, setIsMoveToTrash] = useState(false)
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
                    setIsMoveToTrash(false)
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

    const moveToTrash = (article) => {
        const articleUpdate = {
            title : article.title,
            content : article.content,
            category : article.category,
            status : "thrash",
        }

        fetch(DOMAIN + `/article/${article.id}`, {
            method : "PUT",
            body : JSON.stringify(articleUpdate)
        }).then (res => {return res.json()})
        .then (data => {
            if (data.meta.status === SUCCESS) {
                setIsMoveToTrash(true)
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

    const editArticle = (rowData) => {
        const IDArticle = rowData.id
        navigate(`/article/${IDArticle}`)
    }

    const actionBodyTemplate = (rowData) => {
        return (
            <div className='flex gap-2'>
                <Button
                    icon="pi pi-pencil"
                    className="p-button-rounded p-button-success"
                    onClick={() => editArticle(rowData)}
                    />
                <Button
                    icon="pi pi-trash"
                    className="p-button-rounded p-button-danger"
                    onClick={() => moveToTrash(rowData)}
                />
            </div>
        );
    }

    const leftToolbarTemplate = () => {
        return (
            <div className="body w-[13rem] md:w-full">
                <div className='flex flex-col gap-4 md:flex-row md:gap-2'>
                    <Button
                        label="Published"
                        className="p-button-success mr-2"
                         onClick={() => articleFilter("publish")}
                    />
                    <Button
                        label="Drafts"
                        className="p-button-primary"
                        onClick={() => articleFilter("draft")}
                    />
                    <Button
                        label="Trashed"
                        className="p-button-danger"
                        onClick={() => articleFilter("thrash")}
                    />

                </div >
            </div >
        )
    }

    useEffect(() => {
        getAllArticles()
    }, [,IsMoveToTrash])


    return (<>
        <div className="body">
            <Toolbar
                className="mb-4"
                left={leftToolbarTemplate}
            >
            </Toolbar>

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
                <Column field="title" header="Title" style={{ width: '75%' }} ></Column>
                <Column field="category" header="Category" style={{ width: '50%' }} ></Column>
                <Column header="Action" body={actionBodyTemplate} exportable={false} style={{ minWidth: '8rem' }}></Column>
            </DataTable>
        </div>

    </>);
}

export default AllPosts;