import { ref, isRef, unref, watchEffect } from 'vue'

const baseUrl = "http://localhost:8080/api/v1/"

export function useFetch(url:string) {
    let newUrl:string = baseUrl + url
    const data = ref(null)
    const error = ref(null)
    async function doFetch() {
        data.value = null
        error.value = null
        console.log(newUrl)
        fetch(unref(newUrl))
            .then((res) => res.json())
            .then((json) => (data.value = json))
            .catch((err) => (error.value = err))
        console.log(data.value)
    }

    // setup reactive re-fetch if input URL is a ref
    if (isRef(url))
        watchEffect(doFetch)
    // otherwise, just fetch once and avoid the overhead of a watcher
    else
        doFetch()
    
    return { data, error }
}