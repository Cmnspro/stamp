package mapper

import (
	"stamp/internal/data/dto"
	"stamp/internal/models"
)

func LocalVoteToDTO(vote *models.Vote) dto.Vote {
	return dto.Vote{
		Approved:      vote.Approved,
		DomainStampID: vote.DomainStampID,
		Rating:        vote.Rating,
	}
}

func LocalVoteSliceToDTO(votes models.VoteSlice) dto.VoteSlice {
	res := make(dto.VoteSlice, len(votes))
	for i, vote := range votes {
		res[i] = LocalVoteToDTO(vote)
	}
	return res
}
