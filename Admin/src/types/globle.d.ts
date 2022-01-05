declare interface pictureType {
    CreatedAt: string
    DeletedAt: string | null
    ID: number
    UpdatedAt: string
    category: string | null
    dimension_x: number
    dimension_y: number
    file_size: number
    file_type: string | null
    large: string
    original: string
    path: string | null
    short_url: string | null
    small: string
    url: string | null
    uuid: string
}

declare interface dataType {
    tableData: any,
    centerDialogVisible: boolean,
    banner: pictureType[],
    saveList: pictureType[]
    pictures: number
}

declare interface wallhaven {
    limit: number,
    offset: number
}
declare interface appendBanner {
    (bannerIds: string[]): Promise<{
        code: number,
        message: string
    }>
}

declare interface getBanner {
    (): Promise<{
        code: number,
        message: string,
        result: any
    }>
}