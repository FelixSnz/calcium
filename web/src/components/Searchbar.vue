<template>
   <div class="grow h-18 px-8 pt-2">
      <form autocomplete="off" @submit.prevent="SaveSearchReq()" class="flex items-center">
         <select required v-model="SearchReq.type" class="form-select bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-64  p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
            <option  v-for="(searchType, index) in searchTypes" :key="index"  >{{searchType}}</option>
         </select>
         <label for="voice-search" class="sr-only">Search</label>
         <div class="relative w-full p-2">
            <div class="flex absolute inset-y-0 left-2 items-center pl-3 pointer-events-none">
               <svg class="w-5 h-5 text-gray-500 dark:text-gray-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                  <path fill-rule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clip-rule="evenodd"></path>
               </svg>
            </div>
            <input ref="strkeywords" autocomplete="on" type="text" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full pl-10 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Search by keyword..." required>
         </div>
         <div date-rangepicker datepicker-autohide class="flex w-200">
            <div class="relative grow">
               <div class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
                  <svg aria-hidden="true" class="w-5 h-5 text-gray-500 dark:text-gray-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                     <path fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clip-rule="evenodd"></path>
                  </svg>
               </div>
               <input ref="startdate" name="start" type="text" class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-200 pl-10 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Select date start">
            </div>
            <span class="mx-4 text-gray-500">to</span>
            <div class="relative grow">
               <div class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
                  <svg aria-hidden="true" class="w-5 h-5 text-gray-500 dark:text-gray-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                     <path fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clip-rule="evenodd"></path>
                  </svg>
               </div>
               <input ref="enddate" name="end" type="text" class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-200 pl-10 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Select date end">
            </div>
         </div>
         <button type="submit" class="inline-flex items-center py-2.5 px-3 ml-2 text-sm font-medium text-white bg-blue-700 rounded-lg border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
            <svg class="mr-2 -ml-1 w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
            </svg>
            Search
         </button>
      </form>
      <div v-if="not_found" class="bg-orange-100 border-l-4 border-orange-500 text-orange-700 py-4 pl-4 relative" role="alert">
         <p class="font-bold">Search Failed!</p>
         <p>There where no matches for the keyworkds '{{SearchReq.keywords}}' int the '{{SearchReq.type}}' field and the selected date range</p>
         <span v-on:click="not_found=false" class="absolute top-0 bottom-0 right-0 px-4 py-3">
            <svg class="fill-current h-6 w-6 text-red-500" role="button" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
               <title>Close</title>
               <path d="M14.348 14.849a1.2 1.2 0 0 1-1.697 0L10 11.819l-2.651 3.029a1.2 1.2 0 1 1-1.697-1.697l2.758-3.15-2.759-3.152a1.2 1.2 0 1 1 1.697-1.697L10 8.183l2.651-3.031a1.2 1.2 0 1 1 1.697 1.697l-2.758 3.152 2.758 3.15a1.2 1.2 0 0 1 0 1.698z"/>
            </svg>
         </span>
      </div >
      <div v-if="loading" role="status" class="absolute top-1/3 left-1/2 transform -translate-x-1/2 -translate-y-3/3">
         <svg aria-hidden="true" class="mr-8 w-32 h-32 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/>
            <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"/>
         </svg>
         <span class="sr-only">Loading...</span>
      </div>
      <div v-if="!loading" class="flex-grow">
         <div class="flex py-3">
            <div class="shadow-md sm:rounded-lg h-96 relative  max-w-4xl  max-h-full">
               <table class="overflow-y-auto inline-block min-w-full  align-middle h-80 table-auto grow text-sm text-left text-gray-500 dark:text-gray-400">
                  <thead class="sticky top-0 text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                     <tr>
                        <th scope="col" class="px-4 py-3">Origin</th>
                        <th scope="col" class="px-6 py-3">SubFolder</th>
                        <th scope="col" class="px-10 py-3">Date</th>
                        <th scope="col" class="px-8 py-3">From</th>
                        <th scope="col" class="px-6 py-3">To</th>
                        <th scope="col" class="px-6 py-3">Cc</th>
                     </tr>
                  </thead>
                  <tbody class="overflow-y-auto" v-for="(mail, index) in emails[actual_page]" :key="index">
                     <tr :class="[(active_row == index ? 'active_row' : '') ,'onHover']" v-on:click="HandleEmailClick(mail, index)"  class=" bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                        <td class="px-6 py-4">{{mail.origin}}</td>
                        <td class="px-6 py-4">{{mail.subfolder}}</td>
                        <td class="px-6 py-4">{{new Date(mail.date)}}</td>
                        <td class="px-6 py-4">{{mail.from}}</td>
                        <td v-bind:title="mail.to" class="px-6 py-4">...</td>
                        <td v-bind:title="mail.cc" class="px-6 py-4">...</td>
                     </tr>
                  </tbody>
               </table>
               <div class="absolute bottom-0 left-1/2 transform -translate-x-1/2">
                  <nav aria-label="Page navigation example">
                     <ul class="inline-flex items-center -space-x-px">
                        <li>
                           <a v-on:click="ChangePague(0)" class="block py-2 px-3 ml-0 leading-tight text-gray-500 bg-white rounded-l-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
                              <span class="sr-only">First</span>
                              <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="-40 -40 200 200" xmlns="http://www.w3.org/2000/svg">
                                 <path fill-rule="evenodd" d="M66.6,108.91c1.55,1.63,2.31,3.74,2.28,5.85c-0.03,2.11-0.84,4.2-2.44,5.79l-0.12,0.12c-1.58,1.5-3.6,2.23-5.61,2.2 c-2.01-0.03-4.02-0.82-5.55-2.37C37.5,102.85,20.03,84.9,2.48,67.11c-0.07-0.05-0.13-0.1-0.19-0.16C0.73,65.32-0.03,63.19,0,61.08 c0.03-2.11,0.85-4.21,2.45-5.8l0.27-0.26C20.21,37.47,37.65,19.87,55.17,2.36C56.71,0.82,58.7,0.03,60.71,0 c2.01-0.03,4.03,0.7,5.61,2.21l0.15,0.15c1.57,1.58,2.38,3.66,2.41,5.76c0.03,2.1-0.73,4.22-2.28,5.85L19.38,61.23L66.6,108.91 L66.6,108.91z M118.37,106.91c1.54,1.62,2.29,3.73,2.26,5.83c-0.03,2.11-0.84,4.2-2.44,5.79l-0.12,0.12 c-1.57,1.5-3.6,2.23-5.61,2.21c-2.01-0.03-4.02-0.82-5.55-2.37C89.63,101.2,71.76,84.2,54.24,67.12c-0.07-0.05-0.14-0.11-0.21-0.17 c-1.55-1.63-2.31-3.76-2.28-5.87c0.03-2.11,0.85-4.21,2.45-5.8C71.7,38.33,89.27,21.44,106.8,4.51l0.12-0.13 c1.53-1.54,3.53-2.32,5.54-2.35c2.01-0.03,4.03,0.7,5.61,2.21l0.15,0.15c1.57,1.58,2.38,3.66,2.41,5.76 c0.03,2.1-0.73,4.22-2.28,5.85L71.17,61.23L118.37,106.91L118.37,106.91z" clip-rule="evenodd"></path>
                              </svg>
                           </a>
                        </li>
                        <li>
                           <a v-on:click="ChangePague(actual_page-1)" class="block py-2.5 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
                              <span  class="sr-only">Previous</span>
                              <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                                 <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd"></path>
                              </svg>
                           </a>
                        </li>
                        <li>
                           <a ref="page" aria-current="page" class=" py-2.5 px-3 leading-tight text-blue-600 bg-blue-50 border border-blue-300 hover:bg-blue-100 hover:text-blue-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white">{{actual_page+1}}</a>
                        </li>
                        <li>
                           <a v-on:click="ChangePague(actual_page +1)" class="block py-2.5 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
                              <span  class="sr-only">Next</span>
                              <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                                 <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd"></path>
                              </svg>
                           </a>
                        </li>
                        <li>
                           <a v-on:click="ChangePague(emails.length-1)" class="block py-2 px-3 leading-tight text-gray-500 bg-white rounded-r-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
                              <span class="sr-only">Last</span>
                              <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="-40 -40 200 200" xmlns="http://www.w3.org/2000/svg">
                                 <path fill-rule="evenodd" d="M54.03,108.91c-1.55,1.63-2.31,3.74-2.28,5.85c0.03,2.11,0.84,4.2,2.44,5.79l0.12,0.12c1.58,1.5,3.6,2.23,5.61,2.2 c2.01-0.03,4.01-0.82,5.55-2.37c17.66-17.66,35.13-35.61,52.68-53.4c0.07-0.05,0.13-0.1,0.19-0.16c1.55-1.63,2.31-3.76,2.28-5.87 c-0.03-2.11-0.85-4.21-2.45-5.8l-0.27-0.26C100.43,37.47,82.98,19.87,65.46,2.36C63.93,0.82,61.93,0.03,59.92,0 c-2.01-0.03-4.03,0.7-5.61,2.21l-0.15,0.15c-1.57,1.58-2.38,3.66-2.41,5.76c-0.03,2.1,0.73,4.22,2.28,5.85l47.22,47.27 L54.03,108.91L54.03,108.91z M2.26,106.91c-1.54,1.62-2.29,3.73-2.26,5.83c0.03,2.11,0.84,4.2,2.44,5.79l0.12,0.12 c1.57,1.5,3.6,2.23,5.61,2.21c2.01-0.03,4.02-0.82,5.55-2.37C31.01,101.2,48.87,84.2,66.39,67.12c0.07-0.05,0.14-0.11,0.21-0.17 c1.55-1.63,2.31-3.76,2.28-5.87c-0.03-2.11-0.85-4.21-2.45-5.8C48.94,38.33,31.36,21.44,13.83,4.51l-0.12-0.13 c-1.53-1.54-3.53-2.32-5.54-2.35C6.16,2,4.14,2.73,2.56,4.23L2.41,4.38C0.84,5.96,0.03,8.05,0,10.14c-0.03,2.1,0.73,4.22,2.28,5.85 l47.18,45.24L2.26,106.91L2.26,106.91z" clip-rule="evenodd"></path>
                              </svg>
                           </a>
                        </li>
                     </ul>
                  </nav>
               </div>
            </div>
            <div class="flex grow pl-5 ">
               <div class ="grow shadow-md rounded-lg">
                  <div rows="1" class="py-3 px-5 grow block w-100 min-w-full max-w-xl text-sm text-gray-700 bg-gray-50 rounded-t-lg focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="...">
                     <b>Subject: </b> {{selectedEmail.subject}}
                  </div>
                  <div v-html="selectedEmail.content" style="white-space: pre-line" class="grow overflow-auto max-w-md max-h-80 px-5 block w-100 min-w-full text-sm text-gray-500 bg-white rounded-b-lg focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="...">
                  </div>
               </div>
            </div>
         </div>
      </div>
   </div>
</template>

<script lang="ts">
import type { SearchReq } from '@/interfaces/SearchReq';
import type {EmailData} from '@/interfaces/EmailData'
import SearchEmails from '@/services/SearchService';

//or in file components


export default {

    data() {
        return {
            SearchReq: {} as SearchReq,
            emails: [] as EmailData[][],
            selectedEmail: {} as EmailData,
            active_row: 0,
            not_found:false,
            actual_page: 0,
            loading: false,
            
            searchTypes: [
                "origin",
                "subfolder",
                "subject",
                "from",
                "to",
                "cc",
                "content",
            ]
        }
    },
    methods: {
        async SaveSearchReq() {
         this.emails = [] as EmailData[][]
         this.selectedEmail = {} as EmailData
         this.SearchReq.keywords = (this.$refs.strkeywords as any).value.split(",")
         this.not_found = false

         var from_date = new Date((this.$refs.startdate as any).value);
         var to_date = new Date((this.$refs.enddate as any).value);

         if ((this.$refs.startdate as any).value !="") {
            this.SearchReq.from = from_date.toISOString()
         }
         
         if ((this.$refs.enddate as any).value != "") {
            this.SearchReq.to = to_date.toISOString()
         }
         
         console.log(this.SearchReq.from)
         console.log(this.SearchReq.to)
         this.loading = true
         const res = await SearchEmails(this.SearchReq)
         console.log("response: ", res.data)
         
         if (res.data == null) {
            this.not_found = true
            this.loading = false
            return
         }
         else {

            this.not_found = false
         }

         console.log(res)

         while(res.data.length) this.emails.push(res.data.splice(0,10));
         this.loading = false

        },
         HandleEmailClick (email: EmailData, idx: number) {
            console.log(email)

            this.selectedEmail = email
            this.active_row = idx

            this.SearchReq.keywords.forEach(keyword => {
               const regEx = new RegExp(keyword, "ig");
               const boldKeyword = keyword.toUpperCase().bold()
               this.selectedEmail.content = this.selectedEmail.content.replace(regEx, boldKeyword)
               
               
            });
            

            },
         ChangePague(new_page:number) {
            if (new_page >= 0 && new_page <= this.emails.length-1) {
               this.selectedEmail = {} as EmailData
               this.active_row = 0
               console.log("chaning")
               this.actual_page = new_page;
               console.log(this.emails[this.actual_page])
            }
}
    }}
</script> -->


    
    