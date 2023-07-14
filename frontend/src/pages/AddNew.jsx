import { useState, useEffect } from "react"
import { BAD_REQUEST, CONFLICT, DOMAIN, FAILED, SUCCESS } from "../constant"
import Swal from 'sweetalert2'
import withReactContent from 'sweetalert2-react-content'
import { Textarea, Label, TextInput } from 'flowbite-react';
import { Button } from 'primereact/button';

const AddNew = () => {
    const currentUrl = window.location.href
    const arrCurrentUrl = currentUrl.split('/')
    const id = arrCurrentUrl[4]

    const [status, setStatus] = useState("publish")
    const [isPublish, setIsPublish] = useState(true)
    const [isDraft, setIsDraft] = useState(false)
    const [title, setTitle] = useState('')
    const [content, setContent] = useState('')
    const [category, setCategory] = useState('')
    const MySwal = withReactContent(Swal)

    const setStatusArticle = (status) => {
        if (status === "publish") {
            setIsPublish(true)
            setIsDraft(false)
            setStatus(status)
        } else {
            setIsDraft(true)
            setIsPublish(false)
            setStatus(status)
        }
    }

    const SubmitArticle = (e) => {
        e.preventDefault()
        const newArticle = {
            title: title,
            content: content,
            category: category,
            status: status
        }

        fetch(DOMAIN + '/article', {
            method: "POST",
            body: JSON.stringify(newArticle)
        }).then(res => { return res.json() })
            .then(data => {
                if (data.meta.status === SUCCESS) {
                    MySwal.fire({
                        position: 'center',
                        icon: 'success',
                        title: 'Sukses',
                        text: "Artikel berhasil dibuat",
                        showConfirmButton: true,
                    })
                } else if (data.meta.status === BAD_REQUEST) {
                    MySwal.fire({
                        position: 'center',
                        icon: 'error',
                        title: 'Gagal',
                        text: "inputan tidak sesuai",
                        showConfirmButton: true,
                    })
                } else if (data.meta.status === CONFLICT) {
                    MySwal.fire({
                        position: 'center',
                        icon: 'error',
                        title: 'Gagal',
                        text: `${data.data}`,
                        showConfirmButton: true,
                    })
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


    return (<>
        <div className="form px-10">
            <form className="flex flex-col gap-4" onSubmit={SubmitArticle}>
                <div>
                    <div className="mb-2 block">
                        <Label
                            value="Title"
                        />
                    </div>
                    <TextInput
                        required={true}
                        type="text"
                        onChange={(e) => setTitle(e.target.value)}
                    />
                </div>
                <div>
                    <div className="mb-2 block">
                        <Label
                            value="Content"
                        />
                    </div>
                    <Textarea
                        required={true}
                        rows={10}
                        onChange={(e) => setContent(e.target.value)}
                    />
                </div>
                <div>
                    <div className="mb-2 block">
                        <Label
                            value="Category"
                        />
                    </div>
                    <TextInput
                        required={true}
                        type="text"
                        onChange={(e) => setCategory(e.target.value)}
                    />
                </div>
                <div>
                    <div className="status-button flex flex-col">
                        <div className="label mb-2">
                            <Label
                                value="Content Status (pilih status artikel)"
                            />
                        </div>

                        <div className="button-status flex gap-3">
                            <div className="publish">
                                <Button
                                    type="button"
                                    label="Published"
                                    className={`${isPublish ? "p-button-success" : "p-button-secondary"}`}
                                    onClick={() => setStatusArticle("publish")}
                                />
                            </div>

                            <div className="draft">
                                <Button
                                    type="button"
                                    label="Drafts"
                                    className={`${isDraft ? "p-button-success" : "p-button-secondary"}`}
                                    onClick={() => setStatusArticle("draft")}
                                />
                            </div>
                        </div>
                    </div>
                </div>

                <Button type="submit" className="flex justify-center">
                    <div className="text">
                        Submit
                    </div>
                </Button>
            </form>


        </div>

    </>);
}

export default AddNew;