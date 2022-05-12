import React, { useState, useEffect } from 'react'
import { CollectionSummary } from '../api'
import { Pane, Button, Table, Paragraph, Spinner, majorScale, Pill } from 'evergreen-ui'
import DebouncedFilter from './DebouncedFilter'
import Fuse from 'fuse.js'

type CollectionSummaryFilterProps = {
    summary: CollectionSummary[]
    isLoading: boolean
    onChange(cids: CollectionSummary[]): void
    selected: CollectionSummary[]
}


const fuseOptions: Fuse.IFuseOptions<CollectionSummary> = {
    shouldSort: false,
    keys: ["collectionName"],
    ignoreLocation: true,
    threshold: 0
}

const cleanCollectionName = (name: string) => {
    return name.replace("Hannah Arendt Papers: ", "")
}

const CollectionSummaryFilter = ({ summary, selected, onChange, isLoading }: CollectionSummaryFilterProps) => {
    const [filterText, setFilterText] = useState("")
    const [fuse, setFuse] = useState(new Fuse<CollectionSummary>([], fuseOptions))
    const [filteredSummary, setFilteredSummary] = useState(summary)

    useEffect(() => {
        const buildFuse = () => {
            setFuse(new Fuse(summary, fuseOptions))
        }
        buildFuse()
    }, [summary])

    useEffect(() => {
        const updateFilteredSummary = () => {
            if (filterText) {
                setFilteredSummary(fuse.search(filterText).map(s => s.item))
            } else {
                setFilteredSummary(summary)
            }
        }
        updateFilteredSummary()
    }, [filterText, fuse, summary])


    const handleSelect = (cs: CollectionSummary) => {
        const newSelected = Object.assign([], selected) as CollectionSummary[];
        const index = newSelected.findIndex((s) => s.collectionId === cs.collectionId)
        if (index === -1) {
            newSelected.push(cs)
        } else {
            newSelected.splice(index, 1)
        }

        onChange(newSelected)
    }

    return (
        <Pane width="100%">
            <Table>
                <Table.Head height={32}>
                    <DebouncedFilter onChange={setFilterText} />

                    <Table.TextHeaderCell flexBasis={70} flexShrink={0} flexGrow={0}>results</Table.TextHeaderCell>
                </Table.Head>
                <Table.Head height={32}>
                    <Table.HeaderCell>
                        {selected.length} collections filtered
                    </Table.HeaderCell>
                    <Table.HeaderCell flexBasis={50} flexShrink={0} flexGrow={0}><Button intent="danger" appearance="minimal" disabled={selected.length === 0} onClick={() => onChange([])}>reset filters</Button></Table.HeaderCell>
                </Table.Head>

                {isLoading && <Pane height="50px" display="flex"><Spinner margin="auto" /></Pane>}
                {!isLoading && filteredSummary.length === 0 && <Paragraph padding={majorScale(2)}>No results found</Paragraph>}
                {!isLoading && filteredSummary.length > 0 && <Table.VirtualBody height={240}>
                    {filteredSummary.map((s) => {
                        const isSelected = selected.findIndex(b => b.collectionId === s.collectionId) !== -1
                        return <Table.Row key={`${s.collectionId}${isSelected}`} isSelectable onSelect={() => handleSelect(s)} onDeselect={() => handleSelect(s)} isSelected={isSelected}>
                            <Table.Cell>
                                <Paragraph fontSize={13} lineHeight={'1.3em'}>{cleanCollectionName(s.collectionName)}</Paragraph>
                            </Table.Cell>
                            <Table.TextCell isNumber flexBasis={70} flexShrink={0} flexGrow={0}><Pill>{s.count}</Pill></Table.TextCell>
                        </Table.Row>
                    })}
                </Table.VirtualBody>
                }
            </Table>
        </Pane>
    )
}

export default CollectionSummaryFilter