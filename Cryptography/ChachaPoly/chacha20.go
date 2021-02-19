package ChachaPoly

import (
	"encoding/binary"
	"fmt"
	"math"
)

type ChachaState [16]uint32
type ChachaKey [8]uint32
type ChachaNonce [3]uint32
type ChachaBlockCounter uint32
type ChachaResult [64]uint8
type ChachaEncrypted []uint8

var chachaConstant = [4]uint32{0x61707865, 0x3320646e, 0x79622d32, 0x6b206574}

func main() {
	fmt.Println("Start ChachaPoly")
}

func leftRoll(v uint32, shift int) uint32 {
	overv := v >> (32 - shift)
	v <<= shift
	v += overv
	return v
}

func quarterRound(a, b, c, d uint32) (uint32, uint32, uint32, uint32) {
	a += b
	d ^= a
	d = leftRoll(d, 16)
	c += d
	b ^= c
	b = leftRoll(b, 12)
	a += b
	d ^= a
	d = leftRoll(d, 8)
	c += d
	b ^= c
	b = leftRoll(b, 7)
	return a, b, c, d
}

func (state *ChachaState) Qround(a, b, c, d int) {
	state[a], state[b], state[c], state[d] = quarterRound(state[a], state[b], state[c], state[d])
}

func (state *ChachaState) block() {
	state.Qround(0, 4, 8, 12)
	state.Qround(1, 5, 9, 13)
	state.Qround(2, 6, 10, 14)
	state.Qround(3, 7, 11, 15)
	state.Qround(0, 5, 10, 15)
	state.Qround(1, 6, 11, 12)
	state.Qround(2, 7, 8, 13)
	state.Qround(3, 4, 9, 14)
}

func (counter ChachaBlockCounter) increment(j int) ChachaBlockCounter {
	return ChachaBlockCounter(uint32(counter) + uint32(j))
}

func generateState(key ChachaKey, counter ChachaBlockCounter, nonce ChachaNonce) ChachaState {
	var initState []uint32
	initState = append(initState, chachaConstant[:]...)
	initState = append(initState, key[:]...)
	initState = append(initState, uint32(counter))
	initState = append(initState, nonce[:]...)
	var state ChachaState = [16]uint32{}
	copy(state[:], initState)
	return state
}

func selialize(state ChachaState) ChachaResult {
	result := ChachaResult{}
	for i := range state {
		binary.LittleEndian.PutUint32(result[i*4:(i+1)*4], state[i])
	}
	return result
}

func chacha20Block(key ChachaKey, counter ChachaBlockCounter, nonce ChachaNonce) ChachaResult {
	state := generateState(key, counter, nonce)
	workingState := state
	for i := 0; i < 10; i++ {
		workingState.block()
	}
	for i, v := range workingState {
		state[i] += v
	}
	return selialize(state)
}

func Chacha20Encrypt(key ChachaKey, counter ChachaBlockCounter, nonce ChachaNonce, plainText []byte) ChachaEncrypted {
	var keyStream ChachaResult = ChachaResult{}
	var block ChachaResult = ChachaResult{}
	var encryptedMassage = ChachaEncrypted{}
	for j := 0; j < int(math.Floor(float64(len(plainText)/64))); j++ {
		keyStream = chacha20Block(key, counter.increment(j), nonce)
		copy(block[:], plainText[(j*64):((j+1)*64)])
		for i, v := range keyStream {
			encryptedMassage = append(encryptedMassage, block[i]^v)
		}
	}
	if (len(plainText) % 64) != 0 {
		j := int(math.Floor(float64(len(plainText)) / 64))
		keyStream = chacha20Block(key, counter.increment(j), nonce)
		copy(block[:], plainText[(j*64):])
		for i := 0; i < (len(plainText) % 64); i++ {
			encryptedMassage = append(encryptedMassage, block[i]^keyStream[i])
		}
	}
	return encryptedMassage
}

func Chacha20Decrypt(key ChachaKey, counter ChachaBlockCounter, nonce ChachaNonce, encryptedText []byte) ChachaEncrypted {
	return Chacha20Encrypt(key, counter, nonce, encryptedText)
}
