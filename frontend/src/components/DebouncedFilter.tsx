import React, { useState, useEffect } from 'react'
import { Table } from 'evergreen-ui'

type DebouncedFilterProps = {
    onChange(q: string): void
}

const DebouncedFilter = ({ onChange }: DebouncedFilterProps) => {
    const [filterText, setFilterText] = useState("")
    useEffect(() => {
        const handler = setTimeout(() => onChange(filterText), 500)
        return () => clearTimeout(handler)

    }, [filterText, onChange])
    return <Table.SearchHeaderCell placeholder={"Filter by collection..."} onChange={setFilterText} value={filterText} />
}

export default DebouncedFilter