package torrent 

import (
	"github.com/orion777-cmd/BitTorrent/cmd/bencoding/bencoding"
)

type TorrentFile struct {
	Announce string 
	InfoHash [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length int
	Name int
}

func (bto bencodeTorrent) toTorrentFile() (TorrentFile, error){

}