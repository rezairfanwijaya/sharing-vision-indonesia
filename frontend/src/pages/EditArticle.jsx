import { useState, useEffect } from "react"
import { DOMAIN, SUCCESS } from "../constant"
import Swal from 'sweetalert2'
import withReactContent from 'sweetalert2-react-content'
import { Button, Textarea, Label, TextInput } from 'flowbite-react';

const EditArticle = () => {
    const currentUrl = window.location.href
    const arrCurrentUrl = currentUrl.split('/')
    const id = arrCurrentUrl[4]

    const [article, setArticle] = useState(null)
    const MySwal = withReactContent(Swal)


    const getArticleByID = () => {
        fetch(DOMAIN + `/article/${id}`, {
            method: "GET"
        }).then(res => { return res.json() })
            .then(data => {
                if (data.meta.status === SUCCESS) {
                    setArticle(data.data)
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

    useEffect(() => {
        getArticleByID()
    }, [])


    return (<>
        <div className="form  px-10">
            {article &&
                <form className="flex flex-col gap-4">
                    <div>
                        <div className="mb-2 block">
                            <Label
                                value="Title"
                            />
                        </div>
                        <TextInput
                            value={article.title}
                            required
                            type="text"
                        />
                    </div>
                    <div>
                        <div className="mb-2 block">
                            <Label
                                value="Content"
                            />
                        </div>
                        <Textarea
                            required
                            rows={10}
                            value={article.content}
                        />
                    </div>
                    <div>
                        <div className="mb-2 block">
                            <Label
                                value="Category"
                            />
                        </div>
                        <TextInput
                            value={article.category}
                            required
                            type="text"
                        />
                    </div>

                    <Button type="submit">
                        Submit
                    </Button>
                </form>
            }

        </div>

    </>);
}

export default EditArticle;