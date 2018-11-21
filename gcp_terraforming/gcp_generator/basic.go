package gcp_generator

type BasicGenerator struct{}

type Generator interface {
	Generate(region string) error
}
