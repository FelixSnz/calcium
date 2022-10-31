import axios from './axios'
import type { SearchReq } from '@/interfaces/SearchReq'

 const SearchEmail = async (searchreq: SearchReq) => 
    await axios.post("search_email", searchreq)

export default SearchEmail
