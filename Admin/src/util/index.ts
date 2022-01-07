import dayjs from 'dayjs'
export function qsTime (time:string){ return dayjs(time).format("YYYY-MM-DD-HH-mm")}