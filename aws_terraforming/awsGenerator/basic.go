package awsGenerator

type BasicGenerator struct{}

type Generator interface {
	Generate(region string) error
}
