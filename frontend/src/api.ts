export interface CollectionSummary {
    collectionId: number;
    collectionName: string;
    count: number;
}

export interface SummaryResponse {
    summary: CollectionSummary[]
}

export interface CollectionItem {
    id: number;
    locId: string;
    imageUrl: string;
    OCRResult: string;
    headline: string;
    collectionName: string;
}

export interface SearchResponse {
    results: CollectionItem[]
}

export const fetchSummary = async (query: string): Promise<SummaryResponse> => {
    const res = await fetch("/v1/summary", {
        method: 'POST',
        body: JSON.stringify({ search: query })
    })
    return (res.json() as Promise<SummaryResponse>)
}


export const doSearch = async (query: string, collections: number[], page: number): Promise<SearchResponse> => {
    const res = await fetch("/v1/search", {
        method: 'POST',
        body: JSON.stringify({ search: query, collections, page })
    })
    return (res.json() as Promise<SearchResponse>)
}