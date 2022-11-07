import axios from './axios'
import type { AxiosResponse } from 'axios'
import type { SearchReq } from '@/interfaces/SearchReq'

 const SearchEmails = async (searchreq: SearchReq): Promise<AxiosResponse> => 
    await axios.post("search_mails", searchreq)

export default SearchEmails
