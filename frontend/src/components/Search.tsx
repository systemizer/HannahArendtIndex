import React, { useState } from 'react'
import { SearchInput, majorScale, Pane, Button } from 'evergreen-ui'

type SearchProps = {
    onSubmit(query: string): void
}

const Search = ({ onSubmit }: SearchProps) => {
    const [searchText, setSearchText] = useState<string>("")

    const doSubmit = () => {
        onSubmit(searchText)
    }

    const handleSearchInputChange = (e: React.ChangeEvent<HTMLInputElement>) => setSearchText(e.target.value)

    const handleSearchInputKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key === "Enter") {
            doSubmit()
        }
    }

    return (
        <Pane paddingBottom={majorScale(2)} marginBottom={majorScale(2)} borderBottom display="flex" justifyContent="center" width={"100%"}>
            <SearchInput flex={1} width={'100%'} placeholder={"Search the archives..."} height={majorScale(5)} value={searchText} onKeyDown={handleSearchInputKeyDown} onChange={handleSearchInputChange} />
            <Button marginLeft={majorScale(1)} height={majorScale(5)} onClick={doSubmit}>Search</Button>
        </Pane>
    )
}

export default Search