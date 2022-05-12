import React, { useState } from 'react';
import { Pane, majorScale, Link, Paragraph, Strong, Button } from 'evergreen-ui'
import { CollectionItem } from '../api'
import Highlighter from 'react-highlight-words'

type CollectionItemProps = {
    item: CollectionItem
    searchWords: string[]
}

function getSmallImage(c: CollectionItem): string {
    return c.imageUrl.replace("pct:100", "pct:12.5")
}

const CollectionItemElement = ({ item, searchWords }: CollectionItemProps) => {
    const [showFullText, setShowFullText] = useState(false)

    let disableFullText = false
    if (item.headline.length === item.OCRResult.length) {
        disableFullText = true
    }

    let text = item.headline
    if (showFullText) {
        text = item.OCRResult
    }

    return (

        <Pane backgroundColor={"white"} elevation={1} padding={majorScale(2)} key={item.locId} marginBottom={'20px'}>
            <Pane marginBottom={majorScale(1)} overflow={"hidden"}>
                <Strong>Collection: </Strong><Link target="_blank" href={item.locId}>{item.collectionName}</Link>
            </Pane>
            <Pane display="flex" height={"100%"} overflow={"hidden"}>
                <Pane width={majorScale(20)} marginRight={majorScale(2)} display="flex" flexDirection="column" alignItems="center">
                    <Pane maxHeight={"100px"} overflow={"hidden"}>
                        <Link target="_blank" href={item.locId}>
                            <img alt="Preview" src={getSmallImage(item)} width={"100%"} />
                        </Link>
                    </Pane>
                    <Link size={300} color="neutral" target="_blank" href={item.locId} marginTop={majorScale(1)}>
                        view in archives
                    </Link>
                </Pane>
                <Pane flex={1}>

                    <Pane backgroundColor={"#FAFAFA"} border={"dashed 1px #CCC"} padding={majorScale(1)}>
                        <Paragraph>
                            <Highlighter searchWords={searchWords} textToHighlight={text} />
                        </Paragraph>
                    </Pane>
                    {!disableFullText && !showFullText && <Pane textAlign="center" marginTop={majorScale(1)}>
                        <Button onClick={() => setShowFullText(!showFullText)} appearance="minimal">expand full text...</Button>
                    </Pane>
                    }

                </Pane>
            </Pane>
        </Pane >
    )
}

export default CollectionItemElement