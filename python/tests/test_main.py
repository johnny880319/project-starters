"""Tests for the main module."""

from my_package.main import (
    my_function_1,
    my_function_2,
)


def test_my_function_1() -> None:
    """my_function_1 should return the expected greeting."""
    assert my_function_1() == "Hello, World!"


def test_my_function_2() -> None:
    """my_function_2 should return the expected integer."""
    assert my_function_2() == 42  # noqa: PLR2004
