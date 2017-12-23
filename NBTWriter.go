package GoNBT

import (
	"bytes"
	"compress/gzip"
)

type NBTWriter struct {
	*BinaryStream
}

func NewNBTWriter(network bool, endianType byte) *NBTWriter {
	return &NBTWriter{NewStream([]byte{}, network, endianType)}
}


// WriteUncompressedCompound writes a compound uncompressed.
func (writer *NBTWriter) WriteUncompressedCompound(compound *Compound) {
	writer.PutTag(compound)
	compound.Write(writer)
}


// WriteCompressedCompound writes a compound gzip compressed.
func (writer *NBTWriter) WriteCompressedCompound(compound *Compound) {
	var wr = NewNBTWriter(writer.Network, writer.EndianType)
	wr.WriteUncompressedCompound(compound)

	var buffer = bytes.NewBuffer(wr.GetBuffer())
	var gz = gzip.NewWriter(buffer)
	gz.Write(wr.GetData())
	defer gz.Close()

	writer.PutBytes(buffer.Bytes())
}


// PutTag puts a tag into the buffer.
// This does not yet write payload of the tag.
func (writer *NBTWriter) PutTag(tag INamedTag) {
	writer.PutByte(tag.GetType())
	writer.PutString(tag.GetName())
}


// GetData returns the complete buffer/all data that has been written.
func (writer *NBTWriter) GetData() []byte {
	return writer.Buffer
}